import React from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';

// These pages will still be implemented in the future
import CampaignListPage from './pages/CampaignListPage';
import CampaignFormPage from './pages/CampaignFormPage';

function App () {
  return (
    <BrowserRouter>
      <Routes>
        {/* Campaigns List Page */}
        <Route path="/campaigns" element={<CampaignListPage />} />
        {/* New campaign creation page */}
        <Route path="/campaigns/new" element={<CampaignFormPage />} />
        {/* Edit campaign page */}
        <Route path="/campaigns/:id/edit" element={<CampaignFormPage />} />
        { /* Redirect any other rout to campaigns list */ }
        <Route path="*" element={<Navigate to="/campaigns" replace />} />
      </Routes>
    </BrowserRouter>
  )
};

export default App;