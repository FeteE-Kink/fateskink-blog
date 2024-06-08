import axios from "axios";
import { useAuthStore } from "@/stores/auths";
import { print } from "graphql";

const BASE_URL = `${import.meta.env.VITE_APP_API_BASE_URL}/fateskinkGql`;

const api = axios.create({
  timeout: 5000,
});

api.interceptors.request.use(
  function (config) {
    const authStore = useAuthStore();
    const token = authStore.token;

    if (token) {
      config.headers["Fateskink-Authorization"] = `Bearer ${token}`;
    }

    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

axios.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
  }
);

export default function (
  query,
  variables,
  options = {
    loading: true,
    toast: true,
  }
) {
  // const globalStore = useGlobalStore();
  // globalStore.loading = options.loading;

  return api.post(
    BASE_URL,
    {
      query: print(query),
      variables: variables,
    },
    options
  );
}
