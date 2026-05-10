import { Typography, Stack } from "@mui/material";
import React from "react";

export default function SectionHeader(props: {
  icon: React.ReactNode,
  title: string,
}) {
  return (
    <Stack direction="row" sx={{ alignItems: "center", gap: 1, mb: 2 }} >
      {props.icon}
      <Typography variant="h6" sx={{ fontWeight: 600 }}>{ props.title }</Typography>
    </Stack>
  );
}
