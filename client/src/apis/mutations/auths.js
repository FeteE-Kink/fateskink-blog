import gql from "graphql-tag";

export const SignInMutation = gql`
  mutation SignIn($email: String!, $password: String!) {
    SignIn(email: $email, password: $password) {
      token
    }
  }
`;
