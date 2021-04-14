import { useContext, useEffect, useState } from 'react';
import { SearchContext } from '../Context/SearchContext';
import Items from './Items'
import './Cart.css';

const Cart = () => {
  const { cartList } = useContext(SearchContext)
  const [totAmount, setTotAmount] = useState(0)
  useEffect(() => {
    setTotAmount(0)
    cartList.map((book) => (
      setTotAmount(totAmount => totAmount + book.Price)
    ))
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

      {cartList && <Items books={cartList} />}

    </div>
  );
}

export default Cart;