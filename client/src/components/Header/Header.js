import './Header.css';
import { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';

const Header = () => {
  const [value, setValue] = useState("")
  const handleSearch = (e) => {
    e.preventDefault();
    console.log(value)
    setValue("")
  }
  const history = useHistory()
  const token = localStorage.getItem("token")

  const handleLogOut = () => {
    localStorage.clear("token")
    history.push("/login")
  }

  return (
    <header className="flex">
      <div className="logo">
        <Link to="/">
          <span className="material-icons">
            auto_stories
          </span>
        </Link>
      </div>

      <div className="search-bar flex">
        <form className="search-form flex" onSubmit={handleSearch}>
          <input
            type="search"
            value={value}
            placeholder="Search books..."
            onChange={(e) => setValue(e.target.value)}
          />
          <button type="submit" onClick={handleSearch}>
            <span className="material-icons">
              search
            </span>
          </button>
        </form>

      </div>


      <div className="cart-logo flex">
        <Link to="/cart">
          <span className="material-icons">
            shopping_cart
           </span>
        </Link>
      </div>

      <div className="profile flex">
        <input type="checkbox" className="toggler" id="temp" />

        <span className="material-icons">
          account_circle
          </span>
        <label htmlFor="temp">
          {!token && <ul>
            <Link to="/login"><li>Login</li></Link>
            <Link to="/signup"><li>Sign Up</li></Link>
          </ul>}

          {token && <ul>
            <Link to="/profile"><li>Profile</li></Link>
            <button onClick={handleLogOut}><li>Log Out</li></button>
          </ul>}
        </label>
      </div>
    </header>
  );
}

export default Header;