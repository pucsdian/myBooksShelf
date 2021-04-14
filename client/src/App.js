import './App.css';
import Home from './components/Home/Home';
import Header from './components/Header/Header';
import Login from './components/Login/Login';
import Signup from './components/Signup/Signup';
import Cart from './components/Cart/Cart'
import NotFound from './components/NotFound/NotFound'
import SearchContextProvider from './components/Context/SearchContext'

import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

function App() {
  return (
    <Router>
      <SearchContextProvider>

        <div className="App">

          <div className="content">
            <Switch>
              <Route exact path="/">
                <Header />
                <Home />
              </Route>
              <Route path="/cart">
                <Header />
                <Cart />
              </Route>
              <Route path="/login" component={Login} />
              <Route path="/signup" component={Signup} />
              <Route path="*" component={NotFound} />
            </Switch>
          </div>
        </div>
      </SearchContextProvider>

    </Router>
  );
}

export default App;
