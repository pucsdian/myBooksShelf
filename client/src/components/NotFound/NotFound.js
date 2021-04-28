import { Link } from 'react-router-dom';
import './NotFound.css'

const NotFound = () => {
  return (
    <div className="not-found">
      <h2>404, The page you are looking for is not exist</h2>
      <h2><Link to="/">Click here</Link> to go back to homepage...</h2>
    </div>
  );
}

export default NotFound;