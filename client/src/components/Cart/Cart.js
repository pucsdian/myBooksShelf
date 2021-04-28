import { useContext, useEffect, useState } from 'react';
import { SearchContext } from '../Context/SearchContext';
import Loader from '../Loader/Loader';
import Error from '../Error/Error'
import Items from './Items';
import './Cart.css';

const Cart = () => {
  const { cartList, setCartList } = useContext(SearchContext)
  const [totAmount, setTotAmount] = useState(0)
  const [isPending, setIsPending] = useState(false)
  const [error, setError] = useState(null)

  const removeFromCart = (bookid) => {
    const newCartList = cartList.filter((cart) => cart.ID !== bookid)
    setCartList(newCartList)
  }

  useEffect(() => {
    localStorage.setItem("cartList", JSON.stringify(cartList))
    if (!cartList) {
      setCartList([])
    }
    let amount = 0
    cartList.map((book) => (
      amount += book.Price
    ))
    setTotAmount(amount)
  }, [cartList])


  useEffect(() => {
    setIsPending(true)
    const user = JSON.parse(localStorage.getItem("user"))
    if (!user) {
      setIsPending(false)
      return
    }

    const Authorization = `Bearer ${user.access_token}`

    // headers
    let h = new Headers();
    h.append('Accept', 'application/json');
    h.append('Content-Type', 'application/x-www-form-urlencoded');
    if (Authorization != null) {
      h.append('Authorization', Authorization);
    }
    h.append('Access-Control-Allow-Origin', 'http://localhost:8081')
    h.append('Access-Control-Allow-Headers', 'Authorization')
    h.append('Access-Control-Allow-Methods', 'GET, POST')
    h.append('Access-Control-Expose-Headers', 'Authorization')
    h.append('Host', 'http://localhost:8081');


    const url = 'http://localhost:8081/api/addtocart'

    let req = new Request(url, {
      method: 'GET',
      // mode: 'no-cors',
      headers: h,
    })

    const abortCont = new AbortController();

    fetch(req)
      .then(res => {
        if (!res.ok) { // error coming back from server
          throw Error('could not fetch the data for that resource');
        }
        return res.json();
      })
      .then(data => {
        setIsPending(false);
        setError(null);
        setCartList(data.search_result)
      })
      .catch(err => {
        // setError(err);
        if (err.name === 'AbortError') {
          console.log('fetch aborted')
        } else {
          // auto catches network / connection error
          setIsPending(false);
        }
      })

    // abort the fetch
    return () => abortCont.abort();
  }, [])

  return (
    <div className="cart">
      <div className="cart-header flex">
        <div className="title">
          <h2>You Cart Details</h2>
        </div>
        <div className="amount">
          
          <ul>
            <li><div className="tot-items">Total Items: {cartList.length}</div></li>
            <li><div className="heading">Total Amount: Rs.<span>{totAmount}</span></div></li>
          </ul>
          
          <div className="proceed">
            <button>Proceed To Checkout</button>
          </div>
        </div>
      </div>

      {error && <Error error={error} />}
      {isPending && <div className="pending"><Loader /></div>}
      {cartList.length==0 && <div className="empty-cart"> Your cart is empty</div>}

      {cartList && <Items cartItems={cartList} removeFromCart={removeFromCart} />}

    </div>
  );
}

export default Cart;