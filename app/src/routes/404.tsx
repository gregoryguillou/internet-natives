import { useRouteError } from "react-router-dom";

type ErrorMessage = {
  statusText?: string;
  message?: string;
};

const ErrorPage = () => {
  const error = useRouteError() as ErrorMessage;
  console.error(error);

  return (
    <div id="error-page">
      <h2>Oops!</h2>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{error.statusText || error.message}</i>
      </p>
    </div>
  );
};

export default ErrorPage;
