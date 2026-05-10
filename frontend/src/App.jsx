import { useState, useEffect } from 'react'
import './App.css'
import {
  TextField, Button, Box, Typography, Chip, Card, CardContent, Divider,
  Grid, Table, TableBody, TableCell, TableContainer,
  TableHead, TableRow, Paper, Stack
} from "@mui/material";

import TranslateIcon from "@mui/icons-material/Translate";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import FormatQuoteIcon from "@mui/icons-material/FormatQuote";
import SwapHorizIcon from "@mui/icons-material/SwapHoriz";
import ConjugateIcon from "@mui/icons-material/TableChart";

function SectionHeader({ icon, title }) {
  return (
    <Stack direction="row" sx={{ alignItems: "center", gap: 1, mb: 2 }} >
      {icon}
      <Typography variant="h6" sx={{ fontWeight: 600 }}>{ title }</Typography>
    </Stack>
  );
}

function App() {
  const [input, setInput] = useState('')
  const [result, setResult] = useState(null)

  const handleSearch = async () => {
    const response = await fetch(`http://localhost:8080/word/${input}`)
    const data = await response.json()
    setResult(data)
    console.log("Data:", data)
  }

  useEffect(() => {
    fetch(`http://localhost:8080/word/pulou`)
    	.then(res => res.json())
      .then(data => {
        console.log(data)
        setResult(data)
      })
  }, [])

  return (
    <div>
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

			{result && (
        <Box sx={{ maxWidth: 1000, mx: "auto", px: 3, py: 6 }}>


					{ /* Word and translation */ }

          <Card elevation={0} sx={{ mb: 4, border: "1px solid", borderColor: "divider", borderRadius: 3}}>
            <CardContent sx={{ p: 4 }}>
              <Stack direction={{ xs: "column", sm: "row" }} gap={3} sx={{ alignItems: { sm: "center" }, justifyContent: "space-between" }}>
                <Box>
                  <Typography variant="overline" sx={{ color: "text.secondary", letterSpacing: 2 }}>Português · Brasil</Typography>
                  <Typography variant="h3" sx={{ fontWeight: 700, lineHeight: 1.1 }}>{ result.linguee_search.search_word }</Typography>
                </Box>
                <Divider orientation="vertical" flexItem sx={{ display: { xs: "none", sm: "block" } }} />
                <Box sx={{ textAlign: "right" }}>
                  <Typography variant="overline" sx={{ color: "text.secondary", letterSpacing: 2 }}>Traducción al Español</Typography>
                  <Typography variant="h3" sx={{ fontWeight: 700, color: "primary.main", lineHeight: 1.1 }}>{ result.linguee_search.translation }</Typography>
                </Box>
              </Stack>
            </CardContent>
          </Card>


					{ /* Significados */ }

          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<MenuBookIcon color="primary" />} title="Significados" />

            <Stack sx={{ gap: 1.5 }}>
              { result.dictionary_search.meanings.map((m, i) => (
                <Card key={i} sx={{ border: "1px solid", borderColor: "divider", borderRadius: 2 }}>
                  <CardContent sx={{ py: "12px !important", px: 2.5 }}>
                    <Stack direction="row" sx={{ gap: 2, alignItems: "flex-start" }} >
                      <Typography sx={{ color: "primary.main", fontWeight: 700, minWidth: 24 }}>{ i + 1 }.</Typography>
                      <Typography>{ m }</Typography>
                    </Stack>
                  </CardContent>
                </Card>
              )) }
            </Stack>
          </Box>


					{ /* Ejemplos de uso */ }

					<Box sx={{ mb: 4 }}>
            <SectionHeader icon={<FormatQuoteIcon color="primary" />} title="Ejemplos de uso" />

            <Stack>
              { result.linguee_search.examples.map((ex, i) => (
                <Card key={i} elevation={0} sx={{ border: "1px solid", borderColor: "divider", borderRadius: 2, overflow: "hidden" }}>
                  <Box sx={{ borderLeft: "4px solid", borderColor: "primary.main", px: 2.5, py: 1.5 }}>
                    <Typography sx={{ fontWeight: 500 }}>{ ex.source }</Typography>
                  </Box>
                  <Divider />
                  <Box sx={{ borderLeft: "4px solid", borderColor: "text.disabled", px: 2.5, py: 1.5 }}>
                    <Typography sx={{ color: "text.secondary" }}>{ ex.target }</Typography>
                  </Box>
                </Card>
              )) }
            </Stack>
          </Box>


					{ /* Sinonimos */ }

					<Box sx={{ mb: 4 }}>
            <SectionHeader icon={<SwapHorizIcon color="primary" />} title="Sinónimos" />

            <Box sx={{ display: "flex", flexWrap: "wrap", gap: 1 }}>
              { result.dictionary_search.sinonimos.map((s) => (
                <Chip key={s} label={s} variant="outlined" sx={{ borderRadius: 2 }} />
              )) }
            </Box>
          </Box>


					{ /* Info del verbo */ }

          { result.conjugation_search.verb_info && (
            <Box sx={{ mb: 4 }}>
              <SectionHeader icon={<TranslateIcon color="primary" />} title="Información del verbo" />

              <Card elevation={2} sx={{ border: "1px solid", borderColor: "divider", borderRadius: 2}}>
                <CardContent>
                  <Grid container spacing={2} sx={{ justifyContent: "space-between" }}>
                    <Grid size={{ xs: 6, sm: 3 }} sx={{ textAlign: "center" }}>
                      <Typography variant="caption" sx={{ color: "text.secondary", display: "block", mb: 0.5 }}>Tipo</Typography>
                      <Chip
                        size="small"
                        label={ result.conjugation_search.verb_info.tipo_de_verbo === "regular" ? "Regular" : "Irregular" }
                        color={ result.conjugation_search.verb_info.tipo_de_verbo === "regular" ? "success" : "warning" }
                      />
                    </Grid>
                    {[
                      { label: "Infinitivo", value: result.conjugation_search.verb_info.infinitivo },
                      { label: "Participio", value: result.conjugation_search.verb_info.participio_passado },
                      { label: "Gerúndio", value: result.conjugation_search.verb_info.gerundio },
                    ].map(({ label, value }) => (
                      <Grid key={label} size={{ xs: 6, sm: 3 }} sx={{ textAlign: "center" }}>
                        <Typography variant="caption" sx={{ color: "text.secondary", display: "block", mb: 0.5 }}>{ label }</Typography>
                        <Typography sx={{ fontWeight: 600 }}>{ value }</Typography>
                      </Grid>
                    ))}
                    </Grid>
                  </CardContent>
                </Card>
              </Box>
            ) }


        </Box>
      )}
    </div>
  )
}

export default App
