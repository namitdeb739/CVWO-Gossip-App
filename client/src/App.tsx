import { Box } from '@mantine/core';
import useSWR from "swr";
import './App.css'
import Home from './Pages/Home';

export const ENDPOINT = 'http://localhost:8080'

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  // const {data, mutate} = useSWR(`api/user`, fetcher) 

  return (
    <Home></Home>
  )
}

export default App;
