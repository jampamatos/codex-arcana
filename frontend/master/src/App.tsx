import React, { useEffect, useState } from 'react';

function App () {
  const [version, setVersion] = useState<string | null>(null);

  useEffect(() => {
    // Fetch for /ping endpoint
    fetch('http://localhost:3000/ping')
      .then(res => res.text())
      .then(body => console.log(body))
      .catch(err => console.error('Error fetching ping:', err));

    // Fetch for /api/version endpoint and set version
    fetch('http://localhost:3000/api/version')
      .then(res => res.json())
      .then(data => {
        setVersion(data.version);
      })
      .catch(err => console.error('Error fetching version:', err));
  }, []);
  
  return (
     <div className="bg-blue-500 text-white p-4 text-2xl">
      Hello Codex Arcana!
      {version && (
        <div className="mt-2 text-base">
          <p>Backend Version: <strong>{version}</strong></p>
        </div>
      )}
    </div>
  )
}

export default App;