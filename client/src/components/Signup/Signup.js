import './Signup.css';
import { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import Loader from '../Loader/Loader';

const Signup = () => {
  const [username, setUsername] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [isPending, setIsPending] = useState(false)
  const [error, setError] = useState(null)
  const history = useHistory()

  const handleLogin = (e) => {
    e.preventDefault()

    setIsPending(true)
    const req = { username, password, email }
    fetch('http://localhost:8080/api/signup', {
      method: "POST",
      body: JSON.stringify(req),
      header: {
        "Content-type": "application/json; charset=UTF-8",
        "Access-Control-Allow-Origin": "http://localhost:8080",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE"
      }
    })
      .then(response => response.json())
      .then((data) => {
        setIsPending(false)
        if (data.error) {
          setError(data.error)
        } else {
          history.push("/login")
        }
      }).catch(err => {
        if (err.name === 'AbortError') {
          console.log('fetch aborted')
        } else {
          // auto catches network / connection error
          setIsPending(false);
          setError(err.message);

        }
      })

    setTimeout(() => {
      setError(null)
    }, 3000)

    setUsername("")
    setEmail("")
    setPassword("")

  }


  return (
    <div className="login">
      <div className="my-form">
        <form onSubmit={handleLogin}>

          <div className="logo">
            <Link to="/">
              <span className="material-icons">
                auto_stories
          </span>
            </Link>
          </div>
          <ul>
            <li>{error && <h4>{error}</h4>}</li>
            <li><label >Username</label></li>
            <li><input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value.toLowerCase())}
              required
            /></li>
            <li><label >Email</label></li>
            <li><input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value.toLowerCase())}
              required
            /></li>
            <li><label >Password</label></li>
            <li><input
              type="password"
              name="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            /></li>
            {!isPending && <li><button className="btn btn-background-slide">Sign Up</button></li>}
            {isPending && <li><button className="btn" disabled><Loader /></button></li>}
          </ul>
        </form>
      </div>


      <div className="signup-option">
        Already have an Account? <Link to="/login">Login</Link>
      </div>
    </div>
  );
}

export default Signup;