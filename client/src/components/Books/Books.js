import './Books.css'

const Books = ({ books }) => {
  return (
    <div className="books-list">
      {books.map((book) => (
        <div className="book" key={book.id}>
          <div className="img-section">
            <img src={book.url} alt={`book-${book.id}`} />
          </div>
          <div className="info">
            <h3>{book.title}</h3>
            <h3>{book.author}</h3>
            <h3>{

              Array(Math.round(book.average_rating))
                .fill()
                .map((_) => (
                  <span className="material-icons">
                    star_rate
                  </span>
                ))

            }</h3>
            <h3>Rated By: {book.rating_count}</h3>
            <h3>Price: Rs. {book.price}</h3>
          </div>
          <div className="btn">
            <button> Add To Cart</button>
            <button>Buy Now</button>
          </div>
        </div>
      ))
      }
    </div >
  );
}

export default Books;