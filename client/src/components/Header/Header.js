import './Header.css';
import { useContext, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { SearchContext } from '../Context/SearchContext';


const Header = () => {
  const [searchString, setSearchString] = useState("")
  const { setValue, cartList, setCartList } = useContext(SearchContext)

  var data = JSON.parse(localStorage.getItem("user"))
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
    storeCartList(JSON.parse(localStorage.getItem("user")), JSON.parse(localStorage.getItem("cartList")))
    localStorage.removeItem("user")
    localStorage.removeItem("cartList")
    setCartList([])
    localStorage.clear()
    history.push("/login")
  }

  const bookidList = (cartList) => {
    let temp = []
    cartList.map((book) => (
      temp.push(book.ID)
    ))
    return temp
  }

  const storeCartList = (user, cartList) => {
    console.log(user)
    const bookIDs = bookidList(cartList)
    const Authorization = `Bearer ${user.access_token}`
    const req = { "books_id": bookIDs }
    console.log(req)
    fetch('http://localhost:8081/api/addtocart', {
      mode: 'no-cors',
      method: 'POST',
      body: JSON.stringify(req),
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "*",
        "Access-Control-Allow-Headers": "*",
        "Content-Type": "application/json",
        Authorization
      }
    })
      .then(response => response.json())
      .then((data) => {
        if (data.error) {
          console.log("Cannot update the cart: ", data.error)
        }
        else {
          console.log("Cart is updated: ", data.message)
        }

      }).catch(err => {
        if (err.name === 'AbortError') {
          console.log('fetch aborted')
        } else {
          // auto catches network / connection error
          // setError(err.message)
          console.log("Error: ", err.message)
        }
      })
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
          <div className="tot-items">
           {cartList.length}
          </div>
        </Link>
      </div>

      <div className="profile flex">
        <input type="checkbox" className="toggler" id="temp" />

        <span className="material-icons">
          account_circle
          </span>

        <label htmlFor="temp">
          {!data && <ul>
            <Link to="/login"><li>Login</li></Link>
            <Link to="/signup"><li>Sign Up</li></Link>
          </ul>}

          {data && data.user && <ul>
            <Link to="/profile"><li>Profile</li></Link>
            <button onClick={handleLogOut}><li>Log Out</li></button>
          </ul>}
        </label>
      </div>
      {data && data.user.Username && <div>Welcome <h4>{data.user.Username}</h4></div>}
    </header>

  );
}

export default Header;