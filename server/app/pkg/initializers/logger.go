package initializers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"server/app/pkg/uuid"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CustomTextFormatter formats log entries with each field on a new line
// and highlights field names.
type CustomTextFormatter struct{}

func InitLogger() {
	log.SetFormatter(&CustomTextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// Format renders a single log entry.
func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b bytes.Buffer

	timestamp := entry.Time.Format(time.RFC3339)
	b.WriteString(fmt.Sprintf("\x1b[1m%s\x1b[0m %s\n", entry.Level.String(), timestamp))
	b.WriteString(fmt.Sprintf("Message: %s\n", entry.Message))

	for key, value := range entry.Data {
		b.WriteString(fmt.Sprintf("\x1b[34m%s\x1b[0m: %v\n", key, value))
	}

	files := getStackTrace()
	if files != "" {
		b.WriteString(fmt.Sprintf("Files: \n%s\n", files))
	}

	b.WriteString("\n")
	return b.Bytes(), nil
}

func getStackTrace() string {
	var stackTrace strings.Builder
	pcs := make([]uintptr, 50)
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		projectPaths := []string{
			"server/app/graphql/resolvers",
			"server/app/contrrolers",
			"server/app/services",
		}

		includeFrame := false
		for _, path := range projectPaths {
			if strings.Contains(frame.File, path) {
				includeFrame = true
				break
			}
		}

		if includeFrame {
			stackTrace.WriteString(fmt.Sprintf("%s:%d = %s\n", frame.File, frame.Line, frame.Function))
		}
	}
	return stackTrace.String()
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.RequestURI != "/healthz" {
			handleLogging(ctx)
		}

		startTime := time.Now()
		ctx.Next() // process the request
		endTime := time.Now()

		logWithDuration(ctx, startTime, endTime)
	}
}

func handleLogging(ctx *gin.Context) {
	requestId := uuid.GenerateUUID()
	ctx.Set("request_id", requestId)
	ctx.Writer.Header().Set("X-Request-Id", requestId)

	clientIP := ctx.ClientIP()
	reqMethod := ctx.Request.Method
	headers := filterHeaders(ctx.Request.Header)

	body, _ := io.ReadAll(ctx.Request.Body)
	blurredBody := blurSensitiveData(string(body))

	logRequestDetails(ctx, "Recieved Request", log.Fields{
		"ClientIP": clientIP,
		"Method":   reqMethod,
		"URI":      ctx.Request.RequestURI,
		"Headers":  headers,
		"Body":     prettifyJSON(blurredBody),
	})
}

func logRequestDetails(ctx *gin.Context, message string, fields log.Fields) {
	requestId := ctx.Writer.Header().Get("X-Request-Id")
	fields["RequestID"] = requestId

	log.WithFields(fields).Info(message)
}

func logWithDuration(ctx *gin.Context, start, end time.Time) {
	latencyTime := end.Sub(start)
	statusCode := ctx.Writer.Status()

	logRequestDetails(ctx, "Request Completed", log.Fields{
		"Latency": latencyTime.String(),
		"Status":  statusCode,
	})
}

func prettifyJSON(jsonString string) string {
	var prettifyJSON bytes.Buffer
	err := json.Indent(&prettifyJSON, []byte(jsonString), "", " ")
	if err != nil {
		return jsonString // Return original false string
	}

	return prettifyJSON.String()
}

func filterHeaders(headers map[string][]string) map[string]string {
	filtered := make(map[string]string)
	for key, values := range headers {
		if key != "Cookie" {
			filtered[key] = values[0]
		}
	}

	return filtered
}

func blurSensitiveData(body string) string {
	sensitiveFields := []string{
		"token", "password", "apikey",
	}

	var regexPattern string
	for _, field := range sensitiveFields {
		if regexPattern != "" {
			regexPattern += "|"
		}

		regexPattern += `("` + field + `":")(.*?)(\"|$)`
	}

	re := regexp.MustCompile(regexPattern)

	return re.ReplaceAllString(body, `$1******$3`)
}
