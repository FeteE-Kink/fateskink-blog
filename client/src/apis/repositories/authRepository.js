import api from "@/apis/axios";
import { SignInMutation } from "@/apis/mutations";

export default {
  signIn(input, options = { loading: true, toast: true }) {
    console.log(input);
    return api(SignInMutation, input, options);
  },
};
