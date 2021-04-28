import './Items.css';

const Items = ({ cartItems, removeFromCart }) => {

  return (
    <div className="cart-list">
      {cartItems.map((book) => {
        return (
          <div className="book" key={book.Isbn}>
            <div className="img-section">
              <img src="" alt={book.Isbn} />
              <div><h5>Rs.{book.Price}</h5></div>
            </div>

            <div className="info">
              <div className="line">
                <h5>{book.Title}</h5>
              </div>
              <div className="line">
                <h5>{book.Authors.split('-').join(', ')}</h5>
              </div>
              <div className="line rating">
                <div>   {

                  Array(Math.round(book.AverageRating))
                    .fill()
                    .map((_) => (
                      <span className="material-icons">
                        star_rate
                      </span>
                    ))
                }
                </div>
              </div>
              <div className="line">
                <h5>Rated By: {book.RatingsCount} peoples</h5>
              </div>
            </div>
            <div className="btn">
              <button onClick={() => {
                removeFromCart(book.ID)
              }}> <span className="material-icons">
                  highlight_off
              </span></button>
            </div>
          </div>
        )
      })}

    </div >
  );
}

export default Items;