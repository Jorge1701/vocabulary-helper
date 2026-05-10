import { useState } from "react";
import { Box, Typography, TextField, Button, Grid } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import { SearchResult } from "../model/model";

export default function SearchBar(props: {
  setResults: (results: SearchResult[]) => void
}) {
  const [input, setInput] = useState("")
  const [loading, setLoading] = useState(false)

  const handleSearch = async () => {
    if (input.length > 0 && input.trim() === "") {
      setInput('')
      return
    }
    if (input === undefined || input === "") {
      return
    }

    const words = input.split(/[ ,]/)
    	.filter(v => v.trim() !== "")
      .filter(v => /^[\p{L}-]+$/u.test(v))
      .join(",")

    setLoading(true)

    const response = await fetch(`http://192.168.1.5:8080/words/${words}`)
    const data = await (response.json() as Promise<SearchResult[]>)

    props.setResults(data)
    setLoading(false)
  }

  return (
    <Box sx={{ py: 3, px: 3, display: "flex", flexDirection: "column", alignItems: "center", gap: 5, width: "100%" }}>
      <Box sx={{ textAlign: "center" }}>
        <Typography variant="h2" sx={{ fontWeight: 600 }}>Decifra</Typography>
        <Typography variant="body2" sx={{ color: "text.secondary", mt: 0.5 }}>Busca una o varias palabras para descubrir qué significan</Typography>
      </Box>

      <Grid container sx={{ gap: 0.5, width: "100%", maxWidth: { xs: 400, md: 620 } }}>
        <Grid size={{ xs: 12, md: 8.5 }}>
          <TextField
            fullWidth
            color="info"
            placeholder="ej. pular falamos parabéns"
            size="small"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSearch()}
          />
        </Grid>
        <Grid size={{ xs: 12, md: 3 }}>
          <Button
            fullWidth
            variant="outlined"
            onClick={handleSearch}
            loading={loading}
            loadingPosition="start"
            startIcon={<SearchIcon />}
            sx={{ px: 3, borderRadius: 1, height: "100%" }}
          >
            Buscar
          </Button>
        </Grid>
      </Grid>
    </Box>
  )
}
