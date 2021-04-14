import { createContext, useState } from "react";

export const SearchContext = createContext()

const SearchContextProvider = (props) => {
  const [value, setValue] = useState("")
  const [cartList, setCartList] = useState(JSON.parse(localStorage.getItem("cartList")))
  return (
    <SearchContext.Provider value={{ value, setValue, cartList, setCartList }}>
      {props.children}
    </SearchContext.Provider>
  );
}

export default SearchContextProvider;