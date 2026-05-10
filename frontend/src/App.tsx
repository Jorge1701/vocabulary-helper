import { useState } from 'react'
import "./App.css"
import { Box, Divider } from "@mui/material";

import SearchBar from './components/SearchBar';
import SearchResults from './components/SearchResults';
import WordNotFound from './components/WordNotFound';
import StartMessage from './components/StartMessage';
import { SearchResult } from './model/model';

function App() {
  const [results, setResults] = useState<SearchResult | null>(null)

  return (
    <Box sx={{ maxWidth: 1200, mx: "auto", px: 3, py: 6 }}>

      <SearchBar setResults={setResults} />

      <Divider sx={{ mb: 4 }}/>

      {!results && (
        <StartMessage />
      )}

      {results && results.translation === undefined && (
        <WordNotFound word={results.search_word} />
      )}

			{results && results.translation !== undefined && (
        <SearchResults results={results} />
      )}

    </Box>
  )
}

export default App
