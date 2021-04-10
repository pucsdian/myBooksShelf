import './Home.css'
import Loader from '../Loader/Loader'
import Books from '../Books/Books'
import Error from '../Error/Error'
import useFetch from '../../useFetch';

const Home = () => {
  // const [books, setBooks] = useState([
  //   { title: "Title-1", author: "Shubham-1", url: "https://m.media-amazon.com/images/I/515QCGfqAIL._SL160_.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 1 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780439785969-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 2 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780767923064-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 3 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9781400044733-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 4 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://m.media-amazon.com/images/I/515QCGfqAIL._SL160_.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 14 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780439785969-us.jpg", average_rating: 2.3, rating_count: 33552, price: 250, id: 5 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780767923064-us.jpg", average_rating: 3.26, rating_count: 33552, price: 250, id: 6 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9781400044733-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 7 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://m.media-amazon.com/images/I/515QCGfqAIL._SL160_.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 16 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780439785969-us.jpg", average_rating: 4.2, rating_count: 33552, price: 250, id: 8 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780767923064-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 9 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9781400044733-us.jpg", average_rating: 2.56, rating_count: 33552, price: 250, id: 10 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://m.media-amazon.com/images/I/515QCGfqAIL._SL160_.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 15 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780439785969-us.jpg", average_rating: 1.56, rating_count: 33552, price: 250, id: 11 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9780767923064-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 12 },
  //   { title: "Title-1", author: "Shubham-1", url: "https://pictures.abebooks.com/isbn/9781400044733-us.jpg", average_rating: 3.56, rating_count: 33552, price: 250, id: 13 }
  // ]);


  const { data: books, isPending, error } = useFetch('http://localhost:8080/books')

  return (
    <div className="home">
      {error && <Error error={error} />}
      {isPending && <div className="pending"><Loader /></div>}
      {books && <Books books={books} />}
    </div>
  );
}

export default Home;