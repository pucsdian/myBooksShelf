import { createContext, useState } from "react";

export const SearchContext = createContext()

const SearchContextProvider = (props) => {
  const [value, setValue] = useState("")
  const cartStorage = JSON.parse(localStorage.getItem("cartList"))
  const [cartList, setCartList] = useState(cartStorage ? cartStorage : [])


  return (
    <SearchContext.Provider value={{ value, setValue, cartList, setCartList }}>
      {props.children}
    </SearchContext.Provider>
  );
}

export default SearchContextProvider;