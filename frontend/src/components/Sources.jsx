import { Box, Stack } from "@mui/material";
import LinkIcon from "@mui/icons-material/Link";

import Source from './Source';
import SectionHeader from './SectionHeader';

export default function Sources({ title, sources }) {
  return (
    <Box>
      <SectionHeader icon={<LinkIcon color="primary" />} title={title ?? "Fuentes"} />
  
      <Stack sx={{ gap: 1 }}>
        { sources.map(({label, url}) => (
          <Source key={label} url={url} label={label} />
        )) }
      </Stack>
    </Box>
  )
}
