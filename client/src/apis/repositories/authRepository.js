import api from "@/apis/axios";
import { SignInMutation } from "@/apis/mutations";

export default {
  signIn(input, options = { loading: true, toast: true }) {
    return api(SignInMutation, input, options);
  },
};
