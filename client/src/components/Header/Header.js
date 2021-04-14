import './Header.css';
import { useContext, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { SearchContext } from '../Context/SearchContext';


const Header = () => {
  const [searchString, setSearchString] = useState("")
  const { setValue } = useContext(SearchContext)

  var user = JSON.parse(localStorage.getItem("user"))
  const history = useHistory()


  const handleSearch = (e) => {
    e.preventDefault()
    if (searchString !== "") {
      setValue(searchString)
      setSearchString("")
      history.push("/")
    }
  }

  const handleLogOut = () => {
    localStorage.removeItem("user")
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
            value={searchString}
            placeholder="Search books..."
            onChange={(e) => setSearchString(e.target.value)}
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
          {!user && <ul>
            <Link to="/login"><li>Login</li></Link>
            <Link to="/signup"><li>Sign Up</li></Link>
          </ul>}

          {user && <ul>
            <Link to="/profile"><li>Profile</li></Link>
            <button onClick={handleLogOut}><li>Log Out</li></button>
          </ul>}
        </label>
      </div>
      {user && user.Username && <div>Welcome <h4>{user.Username}</h4></div>}
    </header>

  );
}

export default Header;