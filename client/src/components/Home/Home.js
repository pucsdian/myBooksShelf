import './Home.css'
import Loader from '../Loader/Loader'
import Books from '../Books/Books'
import Error from '../Error/Error'
import { useContext, useEffect, useState } from 'react'
import { SearchContext } from '../Context/SearchContext'

const Home = () => {

  const { value } = useContext(SearchContext)

  const [books, setBooks] = useState(null);
  const [isPending, setIsPending] = useState(true);
  const [error, setError] = useState(null);
  const [initial, setInitial] = useState(false)

  const [totRecords, setTotRecords] = useState(0)
  const [next, setNext] = useState(0)


  useEffect(() => {
    const abortCont = new AbortController();

    fetch('http://localhost:8081/api/randombooks', {
      signal: abortCont.signal,
      header: {
        "Content-type": "application/json; charset=UTF-8",
        "Access-Control-Allow-Origin": "http://localhost:8080",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE"
      }
    })
      .then(res => {
        if (!res.ok) { // error coming back from server
          throw Error('could not fetch the data for that resource');
        }
        return res.json();
      })
      .then(data => {
        setIsPending(false);
        setBooks(data);
        setInitial(true);
        setError(null);
      })
      .catch(err => {
        if (err.name === 'AbortError') {
          console.log('fetch aborted')
        } else {
          // auto catches network / connection error
          setIsPending(false);
          setError(err.message);
        }
      })

    // abort the fetch
    return () => abortCont.abort();
  }, [])





  useEffect(() => {
    setIsPending(true)
    localStorage.removeItem("books")
    setTotRecords(0)

    const req = { "search_string": value }
    fetch('http://localhost:8081/api/books', {
      method: "POST",
      body: JSON.stringify(req),
      header: {
        "Content-type": "application/json; charset=UTF-8",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE"
      }
    })
      .then(response => response.json())
      .then((data) => {
        setIsPending(false)
        if (data.error) {
          setError(data.error)
        }
        else {
          setBooks(data)
          localStorage.setItem("books", JSON.stringify(data))
          if (data && data.total_records) {
            setTotRecords(data.total_records)
          }
          setInitial(false)
        }

      }).catch(err => {
        if (err.name === 'AbortError') {
          console.log('fetch aborted')
        } else {
          // auto catches network / connection error
          setIsPending(false)
          setError(err.message)
        }
      })

  }, [value])

  useEffect(() => {
    setIsPending(true)
    try {
      var records = JSON.parse(localStorage.getItem("books"))
      var start = (next * 52) + 1
      var end = (next + 1) * 52 + 1

      var newPage
      if (records && end >= records.total_records) {
        newPage = records.search_result.slice(start)
      } else if (records) {
        newPage = records.search_result.slice(start, end)
      }
      if (records) {
        records.search_result = newPage
      }
      setIsPending(false)
      setBooks(records)
    } catch (e) {
      console.log(e)
    }

  }, [next])

  const handleDesc = () => {
    setNext(next => next - 1)
  }
  const handleInc = () => {
    setNext(next => next + 1)
  }

  return (
    <div className="home">
      {error && <Error error={error} />}
      {isPending && <div className="pending"><Loader /></div>}
      {books && books.search_result.length === 0 && <div className="result-message"><h3>Sorry, record not found for <span className="red">'{value}'</span>, please search something else..</h3></div>}
      {!initial && books && books.search_result.length !== 0 && <div className="result-message"><h3>Showing results for <span className="green">'{value}'</span>, total {books.total_records} records found</h3></div>}
      <div className="navigation">
        {next >= 1 && <div className="prev"><button onClick={handleDesc}><span className="material-icons">
          arrow_back
</span></button></div>}
        {totRecords > 52 && (next + 1 < totRecords / 52) && <div className="next"><button onClick={handleInc}><span className="material-icons">
          arrow_forward
</span></button></div>}
      </div>
      {books && <Books books={books.search_result} />}
    </div>

  );
}

export default Home;