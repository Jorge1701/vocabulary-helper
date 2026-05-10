import {
  Box, Typography, Chip, Card, CardContent, Divider,
  Grid, Table, TableBody, TableCell, TableContainer,
  TableRow, Stack
} from "@mui/material";

import { SearchResult } from "../model/model";

import TranslateIcon from "@mui/icons-material/Translate";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import FormatQuoteIcon from "@mui/icons-material/FormatQuote";
import SwapHorizIcon from "@mui/icons-material/SwapHoriz";
import ConjugateIcon from "@mui/icons-material/TableChart";

import SectionHeader from "./SectionHeader";
import Sources from "./Sources";

export default function SearchResults(props: {
  results: SearchResult,
}) {

  return (
    <>
        { /* Word and translation */ }

        <Card elevation={0} sx={{ mb: 4, border: "1px solid", borderColor: "divider", borderRadius: 3}}>
          <CardContent sx={{ p: 4 }}>
            <Stack direction={{ xs: "column", sm: "row" }} sx={{ gap: 3, alignItems: { sm: "center" }, justifyContent: "space-between" }}>
              <Box>
                <Typography variant="overline" sx={{ color: "text.secondary", letterSpacing: 2 }}>Português · Brasil</Typography>
                <Typography variant="h3" sx={{ fontWeight: 700, lineHeight: 1.1 }}>{ props.results.found_word }</Typography>
              </Box>
              <Divider orientation="vertical" flexItem sx={{ display: { xs: "none", sm: "block" } }} />
              <Box sx={{ textAlign: "right" }}>
                <Typography variant="overline" sx={{ color: "text.secondary", letterSpacing: 2 }}>Traducción al Español</Typography>
                <Typography variant="h3" sx={{ fontWeight: 700, color: "primary.main", lineHeight: 1.1 }}>{ props.results.translation }</Typography>
              </Box>
            </Stack>
          </CardContent>
        </Card>


        { /* Significados */ }


				{ props.results.meanings && (
          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<MenuBookIcon color="primary" />} title="Significados" />

            <Stack sx={{ gap: 1.5 }}>
              { props.results.meanings.map((m, i) => (
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
        ) }


        { /* Ejemplos de uso */ }

        { props.results.examples && (
          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<FormatQuoteIcon color="primary" />} title="Ejemplos de uso" />

            <Stack>
              { props.results.examples.map((ex, i) => (
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

				{ props.results.synonyms && (
          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<SwapHorizIcon color="primary" />} title="Sinónimos" />

            <Box sx={{ display: "flex", flexWrap: "wrap", gap: 1 }}>
              { props.results.synonyms.map((s) => (
                <Chip key={s} label={s} variant="outlined" sx={{ borderRadius: 2 }} />
              )) }
            </Box>
          </Box>
        ) }


        { /* Info del verbo */ }

        { props.results.verb_info && (
          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<TranslateIcon color="primary" />} title="Información del verbo" />

            <Card elevation={2} sx={{ border: "1px solid", borderColor: "divider", borderRadius: 2}}>
              <CardContent>
                <Grid container spacing={2} sx={{ justifyContent: "space-between" }}>
                  <Grid size={{ xs: 6, sm: 3 }} sx={{ textAlign: "center" }}>
                    <Typography variant="caption" sx={{ color: "text.secondary", display: "block", mb: 0.5 }}>Tipo</Typography>
                    <Chip
                      size="small"
                      label={ props.results.verb_info.type === "regular" ? "Regular" : "Irregular" }
                      color={ props.results.verb_info.type === "regular" ? "success" : "warning" }
                    />
                  </Grid>
                  {[
                    { label: "Infinitivo", value: props.results.verb_info.infinitive },
                    { label: "Participio", value: props.results.verb_info.past_participle },
                    { label: "Gerúndio", value: props.results.verb_info.present_participle },
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

        { props.results.verb_info && (
          <Box sx={{ mb: 4 }}>
            <SectionHeader icon={<ConjugateIcon color="primary" />} title="Conjugaciones" />

            <Grid container spacing={2}>
              {[
                { label: "Presente", values: props.results.verb_info.simple_present },
                { label: "Pretérito Imperfeito", values: props.results.verb_info.imperfect_past },
                { label: "Pretérito Perfeito", values: props.results.verb_info.simple_past },
                { label: "Pretérito Mais-que-perfeito", values: props.results.verb_info.perfect_past },
                { label: "Futuro do Presente", values: props.results.verb_info.simple_future },
                { label: "Futuro do Pretérito", values: props.results.verb_info.conditional },
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

				{ props.results.sources && (
					<Sources title="Fuentes" sources={props.results.sources}/>
        ) }

    </>
  )
}
