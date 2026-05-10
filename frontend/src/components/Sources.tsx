import { Box, Stack } from "@mui/material";
import LinkIcon from "@mui/icons-material/Link";

import Source from "./Source";
import SectionHeader from "./SectionHeader";

export default function Sources(props: {
  title: string,
  sources: Record<string, string>,
}) {
  return (
    <Box>
      <SectionHeader icon={<LinkIcon color="primary" />} title={props.title} />
  
      <Stack sx={{ gap: 1 }}>
        { Object.entries(props.sources).map(([label, url]) => (
          <Source key={label} url={url} label={label} />
        )) }
      </Stack>
    </Box>
  )
}
