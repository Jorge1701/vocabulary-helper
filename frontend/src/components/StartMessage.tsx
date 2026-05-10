import { Box, Typography } from "@mui/material";
import MenuBookIcon from "@mui/icons-material/MenuBook";

export default function StartMessage() {
  return (
    <Box sx={{ flex: 1, display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center", gap: 1.5, opacity: 0.4, py: 0 }}>
      <MenuBookIcon sx={{ fontSize: 52 }}/>
      <Typography variant="body2" sx={{ color: "text.secondary", textAlign: "center" }}>Ingresa una o varias palabras para comenzar</Typography>
    </Box>
  )
}
