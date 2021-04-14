import './Login.css';
import { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import Loader from '../Loader/Loader';

const Login = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [isPending, setIsPending] = useState(false)
  const [error, setError] = useState(null)
  const history = useHistory()

  const handleLogin = (e) => {
    e.preventDefault()
    setIsPending(true)
    const req = { "username": email, password }
    fetch('http://localhost:8080/api/login', {
      method: "POST",
      body: JSON.stringify(req),
      header: {
        "Content-type": "application/json; charset=UTF-8",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE"
      }
    })
      .then(response => response.json())
      .then((data) => {
        setIsPending(false)
        if (data && data.error) {
          setError(data.error)
        }
        else {
          try {
            localStorage.setItem("user", JSON.stringify(data.user))
          } catch (e) {
            console.log(e)
          }

          history.push("/")
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
            <li><label >Username or email address</label></li>
            <li><input
              type="text"
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
            {!isPending && <li><button className="btn btn-background-slide">Log In</button></li>}
            {isPending && <li><button className="btn" disabled><Loader /></button></li>}
            <li><p>Forgot Password? <Link to="/forgot-password">Click Here</Link></p></li>
          </ul>
        </form>
      </div>


      <div className="signup-option">
        Need an Account? <Link to="/signup">Sign Up</Link>
      </div>
    </div>
  );
}

export default Login;