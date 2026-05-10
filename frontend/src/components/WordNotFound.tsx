import { Box, Typography, Card, CardContent } from "@mui/material";
import SearchOffIcon from "@mui/icons-material/SearchOff";

import Sources from "./Sources";

export default function WordNotFound(props: {
  word: string,
}) {
  return (
    <Box>

      <Box sx={{ textAlign: "center", mb: 5 }}>
        <SearchOffIcon sx={{ fontSize: 52, color: "text.secondary", opacity: 0.4, mb: 2 }} />
        <Typography variant="h6" sx={{ fontWeight: 600, mb: 1 }}>
          No se encontraron resultados para{" "}
          <Typography component="span" sx={{ fontWeight: 700, color: "primary.main" }}>"{props.word}"</Typography>
        </Typography>
        <Typography variant="body2" sx={{ color: "text.secondary" }}>Verificá la ortografía o intentá con otra palabra.</Typography>
      </Box>

      <Sources title="También puedes probar:" sources={{
        "Reverso context": `https://context.reverso.net/traduccion/portugues-espanol/${props.word}`,
        "Dicio": `https://www.dicio.com.br/pesquisa.php?q=${props.word}` ,
        "Conjugacao": `https://www.conjugacao.com.br/busca.php?q=${props.word}` ,
        "Linguee" : `https://www.linguee.es/espanol-portugues/search?query=${props.word}` 
      }}/>

    </Box>
  );
}
