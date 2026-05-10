import { Typography, Stack } from "@mui/material";

export default function SectionHeader({ icon, title }) {
  return (
    <Stack direction="row" sx={{ alignItems: "center", gap: 1, mb: 2 }} >
      {icon}
      <Typography variant="h6" sx={{ fontWeight: 600 }}>{ title }</Typography>
    </Stack>
  );
}
