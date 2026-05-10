import React, { useState } from 'react'
import "./App.css"
import { Box, Divider, Tab, Tabs } from "@mui/material";

import SearchBar from './components/SearchBar';
import SearchResults from './components/SearchResults';
import WordNotFound from './components/WordNotFound';
import StartMessage from './components/StartMessage';
import { SearchResult } from './model/model';

function TabPanel(props: {
  children: React.ReactNode,
  index: number,
  currentTab: number,
}) {
  return (
    <div
    	role="tabpanel"
      hidden={ props.currentTab !== props.index }>
      { props.currentTab === props.index && (
        <Box>
          { props.children }
        </Box>
      ) }
    </div>
  )
}

function App() {
  const [results, setResults] = useState<SearchResult[] | null>(null)
  const [currentTab, setCurrentTab] = useState(0)

  return (
    <Box sx={{ maxWidth: 1200, mx: "auto", px: 3, py: 6 }}>

      <SearchBar setResults={setResults} />

      <Divider sx={{ mb: 4 }}/>

			{ results ? (
        <>
          <Tabs
            orientation="horizontal"
            variant="scrollable"
            value={currentTab}
            onChange={(_, newTab) => setCurrentTab(newTab) }>
              { results.map(({search_word}) => (
              	<Tab key={search_word} label={search_word} />
              )) }
          </Tabs>
          { results.map((result, i) => (
            <TabPanel key={result.search_word} currentTab={currentTab} index={i}>
              {result && result.translation !== undefined ? (
                <SearchResults results={result} />
              ) : (
                <WordNotFound word={result.search_word} />
              )}
            </TabPanel>
          )) }
        </>
      ) : (
        <StartMessage />
      )}

    </Box>
  )
}

export default App
