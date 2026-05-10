import { useState } from "react";
import { Box, Typography, TextField, Button, Divider, InputAdornment } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";

export default function SearchBar({ setResults }) {
  const [ input, setInput ] = useState('')
  const [ loading, setLoading ] = useState(false)

  const handleSearch = async () => {
    if (input.length > 0 && input.trim() === '') {
      setInput('')
      return
    }
    if (input === undefined || input === '') {
      return
    }
    setLoading(true)

    const response = await fetch(`http://localhost:8080/word/${input}`)
    const data = await response.json()
    setResults({
      input_word: input,
      found_word: data.linguee_search.search_word,
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

    setLoading(false)
  }

  return (
    <Box sx={{ py: 6, px: 3, display: "flex", flexDirection: "column", alignItems: "center", gap: 5, width: "100%" }}>
      <Box sx={{ textAlign: "center" }}>
        <Typography variant="h2" sx={{ fontWeight: 600 }}>Decifra</Typography>
        <Typography variant="body2" sx={{ color: "text.secondary", mt: 0.5 }}>Descubrí qué significa cada palabra</Typography>
      </Box>

      <Box sx={{ display: "flex", gap: 1.5, width: "100%", maxWidth: 520 }}>
        <TextField
          fullWidth
          color="info"
          placeholder="Busca una palabra... ej. pular"
          size="small"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && handleSearch()}
        />
        <Button
        	variant="outlined"
          size="small"
          onClick={handleSearch}
          loading={loading}
          loadingPosition="start"
          startIcon={<SearchIcon />}
          sx={{ px: 3, borderRadius: 2 }}
        >
          Buscar
        </Button>
      </Box>
    </Box>
  )
}
