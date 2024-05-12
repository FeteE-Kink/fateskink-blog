import { ref } from "vue";

import { defineStore } from "pinia";
import { get } from "lodash";

import axios from "@/apis/axios";

export const useGlobalStore = defineStore("global", () => {
  const validationErrors = ref({});
  const errorMessage = ref("");

  function getErrors(name) {
    return get(validationErrors.value, name, []);
  }

  function setErrorMessage(message) {
    errorMessage.value = message;
  }

  function setValidationErrors(value) {
    validationErrors.value = value;
  }

  async function uploadImages(formData) {
    try {
      const { data } = await axios.post("/upload", formData);
      return data.data;
    } catch (error) {
      alert(error);
    }
  }

  return {
    // Errors Msgs
    errorMessage,
    validationErrors,

    uploadImages,
    setValidationErrors,
    getErrors,
    setErrorMessage,
  };
});
