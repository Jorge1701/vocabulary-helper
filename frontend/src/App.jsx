import { useState } from 'react'
import './App.css'
import {
  TextField, Button, Box, Typography, Chip, Card, CardContent, Divider,
  Grid, Table, TableBody, TableCell, TableContainer,
  TableRow, Stack, Link
} from "@mui/material";

import TranslateIcon from "@mui/icons-material/Translate";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import FormatQuoteIcon from "@mui/icons-material/FormatQuote";
import SwapHorizIcon from "@mui/icons-material/SwapHoriz";
import ConjugateIcon from "@mui/icons-material/TableChart";
import LinkIcon from "@mui/icons-material/Link";

function FaviconImg({ url }) {
  const domain = new URL(url).hostname;
  return (
    <img
      src={`https://www.google.com/s2/favicons?domain=${domain}&sz=32`}
      width={20}
      height={20}
      style={{ borderRadius: 4 }}
    />
  );
}

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
        <Box sx={{ maxWidth: 1200, mx: "auto", px: 3, py: 6 }}>


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

					{ result.linguee_search.examples && (
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
          ) }


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

          { result.conjugation_search.found && (
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


					{ /* Tablas de conjugaciones */ }

          { result.conjugation_search.found && (
            <Box sx={{ mb: 4 }}>
              <SectionHeader icon={<ConjugateIcon color="primary" />} title="Conjugaciones" />

							<Grid container spacing={2}>
                {[
                  { label: "Presente", values: result.conjugation_search.presente },
                  { label: "Pretérito Imperfeito", values: result.conjugation_search.preterito_imperfeito },
                  { label: "Pretérito Perfeito", values: result.conjugation_search.preterito_perfeito },
                  { label: "Pretérito Mais-que-perfeito", values: result.conjugation_search.preterito_mais_que_perfeito },
                  { label: "Futuro do Presente", values: result.conjugation_search.futuro_do_presente },
                  { label: "Futuro do Pretérito", values: result.conjugation_search.futuro_do_preterito },
                ].map(({ label, values }) => (
                  <Grid size={{ xs: 12, sm: 6, md: 4 }} key={label}>
                    <Card sx={{ border: "1px solid", borderColor: "divider", borderRadius: 2, height: "100%" }}>
                      <CardContent sx={{ pb: "12px !important" }}>
                        <Typography variant="subtitle2" sx={{ fontWeight: 700, mb: 1.5, color: "primary.main" }}>{ label }</Typography>
                        <TableContainer>
                          <Table size="small">
                            <TableBody>
                              {[
                                { person: "Eu", value: values.first_per_sin },
                                { person: "Tu", value: values.second_per_sin },
                                { person: "Ele/ela/você", value: values.third_per_sin },
                                { person: "Nós", value: values.first_per_plu },
                                { person: "Vós", value: values.second_per_plu },
                                { person: "Eles/elas/vocês", value: values.third_per_plu },
                              ].map(({ person, value }) => (
                                <TableRow key={person} sx={{ "&:last-child td": { border: 0 } }}>
                                  <TableCell sx={{ color: "text.secondary", fontSize: "0.78rem", pl: 0, width: "45%" }}>{ person }</TableCell>
                                  <TableCell sx={{ fontWeight: 500, pr: 0 }}>{ value }</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>
                      </CardContent>
                    </Card>
                  </Grid>
                ))}
              </Grid>
            </Box>
          ) }


					{ /* Sources */ }

          <Box>
            <SectionHeader icon={<LinkIcon color="primary" />} title="Fuentes" />
        
            <Stack sx={{ gap: 1 }}>
              {[
                { label: "Dicio", url: result.dictionary_search.source_url },
                { label: "Conjugacao", url: result.conjugation_search.source_url },
                { label: "Linguee", url: result.linguee_search.source_url },
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
        </Box>
      )}
    </div>
  )
}

export default App
