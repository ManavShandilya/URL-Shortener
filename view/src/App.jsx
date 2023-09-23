import { useState } from "react"
import Heading from "./components/Heading"
import ShortenResponse from "./components/ShortenResponse"
import URLinput from "./components/URLinput"

function App() {

  const [response, setresponse] = useState(null)

  return (
    <div className="container mx-auto mt-24 p-4">
      <Heading />
      <URLinput response={response} setresponse={setresponse} />
      <ShortenResponse response={response} />
    </div>
  )
}

export default App
