import './Error.css'

const Error = ({ error }) => {
  return (
    <div className="errors">
      <div className="message">
        <h2>{error}</h2>
        <h4>Please try again later</h4>
      </div>
    </div>

  );
}

export default Error;