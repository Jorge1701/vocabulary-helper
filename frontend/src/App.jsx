import { useState } from 'react'
import './App.css'
import {
  TextField, Button, Box, Stack, Typography, Card, CardContent, Link
} from "@mui/material";
import LinkIcon from "@mui/icons-material/Link";

import SearchResults from './components/SearchResults';
import SectionHeader from './components/SectionHeader';
import FaviconImg from './components/FaviconImg';

function App() {
  const [input, setInput] = useState('')
  const [results, setResults] = useState(null)

  const handleSearch = async () => {
    const response = await fetch(`http://localhost:8080/word/${input}`)
    const data = await response.json()
    setResults({
      search_word: data.linguee_search.search_word,
      translation: data.linguee_search.translation,
      meanings: data.dictionary_search.meanings,
      examples: data.linguee_search.examples,
      synonyms: data.dictionary_search.sinonimos,
      verb_info: data.conjugation_search.found ? data.conjugation_search.verb_info : undefined,
      conjugations: data.conjugation_search.found ? data.conjugation_search : undefined,
      sources: {
        dicio: data.dictionary_search.source_url,
        conjugacao: data.conjugation_search.source_url,
        linguee: data.linguee_search.source_url,
      }
    })
  }

  return (
    <Box sx={{ maxWidth: 1200, mx: "auto", px: 3, py: 6 }}>
      <Box component="section" sx={{ p: 2, border: '1px dashed grey' }}>
      	<TextField
        	label="Palabra"
          variant="outlined"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />

      	<Button variant="contained" onClick={handleSearch}>
          Buscar
        </Button>
      </Box>

			{results && (
        <>
          { /* Results */ }

          <SearchResults results={results} />

          { /* Sources */ }

          <Box>
            <SectionHeader icon={<LinkIcon color="primary" />} title="Fuentes" />
        
            <Stack sx={{ gap: 1 }}>
              {[
                { label: "Dicio", url: results.sources.dicio },
                { label: "Conjugacao", url: results.sources.conjugacao },
                { label: "Linguee", url: results.sources.linguee },
                { label: "Reverso context", url: `https://context.reverso.net/traduccion/portugues-espanol/${input}` },
              ].filter(item => item.url !== undefined).map(({label, url}) => (
                <Card key={label} elevation={1} sx={{ borderRadius: 2 }}>
                  <CardContent sx={{ py: "10px !important", px: 2 }}>
                    <Stack direction="row" sx={{ alignItems: "center", justifyContent: "space-between" }}>
                      <Stack direction="row" sx={{ alignItems: "center", gap: 1.5 }}>
                        <FaviconImg url={url} />
                        <Typography variant="body2" sx={{ color: "text.secondary" }}>{ label }</Typography>
                      </Stack>
                      <Link variant="body2" href={url} target="_blank" rel="noopener noreferrer">{ url }</Link>
                    </Stack>
                  </CardContent>
                </Card>
              ))}
            </Stack>
          </Box>
        </>
      )}
    </Box>
  )
}

export default App
