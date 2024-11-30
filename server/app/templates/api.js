export async function executeApi(query, variables) {
  const response = await fetch(`${BASE_URL}fateskinkGql`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ query: query, variables: variables }), // Gửi query và variables
  });
}
