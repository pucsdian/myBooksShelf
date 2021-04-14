import { useContext } from 'react';
import { SearchContext } from '../Context/SearchContext';
import './Books.css'

const Books = ({ books }) => {
  const { cartList, setCartList } = useContext(SearchContext)


  const addToCart = (bookid, e) => {
    e.target.innerHTML = "Added to Cart"
    e.target.setAttribute("disabled", true)
    e.target.setAttribute('style', 'cursor:not-allowed')

    const book = books.filter(book => book.ID === bookid)
    setCartList([...cartList, book[0]])

    setTimeout(() => {
      e.target.innerHTML = "Add to Cart"
      e.target.setAttribute("disabled", false)
      e.target.setAttribute('style', 'cursor:pointer')
    }, 3000)

  }

  return (
    <div className="books-list">
      {books.map((book) => {
        return (
          <div className="book" key={book.ID}>
            <div className="img-section">
              <img src="" alt={`isbn-${book.Isbn}`} />
            </div>
            <div className="info">
              <div className="line">
                <h4>Title: </h4>
                <h5>{book.Title}</h5>
              </div>
              <div className="line">
                <h4>Authors:</h4>
                <h5>{book.Authors.split('-').join(', ')}</h5>
              </div>
              <div className="line rating">
                <div><h4>Rating:</h4></div>
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
              <button onClick={(e) => {
                addToCart(book.ID, e)
              }}> Add to Cart</button>
              <button>Buy @ Rs.{book.Price}</button>
            </div>
          </div>
        )
      })}

    </div >
  );
}

export default Books;