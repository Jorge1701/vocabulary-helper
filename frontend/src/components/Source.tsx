import { Stack, Typography, Card, CardContent, Link } from "@mui/material";

import FaviconImg from "./FaviconImg";

export default function Source(props: {
  url: string,
  label: string,
}) {
  return (
    <Card elevation={1} sx={{ borderRadius: 2 }}>
      <CardContent sx={{ py: "10px !important", px: 2 }}>
        <Stack direction="row" sx={{ alignItems: "center", justifyContent: "space-between" }}>
          <Stack direction="row" sx={{ alignItems: "center", gap: 1.5 }}>
            <FaviconImg url={props.url} />
            <Typography variant="body2" sx={{ color: "text.secondary" }}>{ props.label }</Typography>
          </Stack>
          <Link variant="body2" href={props.url} target="_blank" rel="noopener noreferrer">{ props.url }</Link>
        </Stack>
      </CardContent>
    </Card>
  )
}
