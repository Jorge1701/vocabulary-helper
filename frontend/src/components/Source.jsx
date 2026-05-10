import { Stack, Typography, Card, CardContent, Link } from "@mui/material";

import FaviconImg from './FaviconImg';

export default function Source({ url, label }) {
  return (
    <Card elevation={1} sx={{ borderRadius: 2 }}>
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
  )
}
