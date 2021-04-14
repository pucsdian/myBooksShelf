import { useContext, useEffect, useState } from 'react';
import { SearchContext } from '../Context/SearchContext';
import Items from './Items'
import './Cart.css';

const Cart = () => {
  const { cartList, setCartList } = useContext(SearchContext)
  const [totAmount, setTotAmount] = useState(0)

  const removeFromCart = (bookid) => {
    const newCartList = cartList.filter((cart) => cart.ID !== bookid)
    setCartList(newCartList)
  }

  useEffect(() => {
    localStorage.setItem("cartList", JSON.stringify(cartList))
    let amount = 0
    cartList.map((book) => (
      amount += book.Price
    ))
    setTotAmount(amount)
  }, [cartList])


  return (
    <div className="cart">
      <div className="cart-header flex">
        <div className="title">
          <h2>You Cart Details</h2>
        </div>
        <div className="amount">
          <div className="heading">Total Amount: Rs.<span>{totAmount}</span></div>
          <div className="proceed">
            <button>Proceed To Checkout</button>
          </div>
        </div>
      </div>

      {cartList && <Items cartItems={cartList} removeFromCart={removeFromCart} />}

    </div>
  );
}

export default Cart;