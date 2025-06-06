import React, { useEffect, useState } from 'react';
import type { FormEvent } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

// Interface that reflects backend Campaign model
interface Campaign {
  id?: number;
  name: string;
  description: string;
  created_at?: string;
  updated_at?: string;
}

const CampaignFormPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const isEditMode = Boolean(id);
  const navigate = useNavigate();

  const [campaign, setCampaign] = useState<Campaign>({
    name: '',
    description: '',
  });
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // If the user is trying to access the edit page without an ID, redirect to the list
  useEffect(() => {
    if (!isEditMode) return;

    const fetchCampaign = async () => {
      try {
        setLoading(true);
        setError(null);

        const response = await fetch(`http://localhost:3000/api/campaigns/${id}`);
        if (response.status === 404) {
          alert('Campaing not found.');
          navigate('/campaigns');
          return;
        }
        if (!response.ok) {
          throw new Error(`Failed to fetch campaign: ${response.statusText}`);
        }

        const data: Campaign = await response.json();
        setCampaign({
          name: data.name,
          description: data.description,
        });
      } catch (err: unknown) {
        console.error(err);
        if (err instanceof Error) {
          setError(err.message || 'Unknown error');
        } else {
          setError('Unknown error');
        }
      } finally {
        setLoading(false);
      }
    };

    fetchCampaign();
  }, [id, isEditMode, navigate]);

  // Function to submit data to the backend (POST or PUT)
  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (campaign.name.trim() === '') {
      setError("The 'Name' field is required.");
      return;
    }

    try {
      setLoading(true);
      setError(null);

      const url = isEditMode
        ? `http://localhost:3000/api/campaigns/${id}`
        : 'http://localhost:3000/api/campaigns';
      const method = isEditMode ? 'PUT' : 'POST';

      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: campaign.name,
          description: campaign.description,
        }),
      });

      if (response.status === 400) {
        const text = await response.text();
        throw new Error(text || 'Validation error');
      }
      if (!response.ok) {
        throw new Error(`Failed to save: ${response.statusText}`);
      }

      // Upon saving, redirect to the campaigns list
      navigate('/campaigns');
    } catch (err: unknown) {
      console.error(err);
      if (err instanceof Error) {
        setError(err.message || 'Failed to save campaign');
      } else {
        setError('Failed to save campaign');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-4 max-w-lg mx-auto">
      <h1 className="text-2xl font-bold mb-4">
        {isEditMode ? 'Edit Campaign' : 'New Campaign'}
      </h1>

      {error && (
        <div className="mb-4 text-red-500">
          <p>{error}</p>
        </div>
      )}

      {loading ? (
        <p>Loading...</p>
      ) : (
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="name" className="block mb-1 font-medium">
              Name<span className="text-red-500">*</span>
            </label>
            <input
              id="name"
              type="text"
              value={campaign.name}
              onChange={e => setCampaign(prev => ({ ...prev, name: e.target.value }))}
              className="w-full border border-gray-300 rounded px-3 py-2"
              required
            />
          </div>

          <div>
            <label htmlFor="description" className="block mb-1 font-medium">
              Description
            </label>
            <textarea
              id="description"
              value={campaign.description}
              onChange={e =>
                setCampaign(prev => ({ ...prev, description: e.target.value }))
              }
              className="w-full border border-gray-300 rounded px-3 py-2"
              rows={4}
            />
          </div>

          <div className="flex justify-end space-x-2">
            <button
              type="button"
              onClick={() => navigate('/campaigns')}
              className="bg-gray-300 hover:bg-gray-400 text-gray-800 px-4 py-2 rounded"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={loading}
              className={`${
                loading ? 'bg-blue-300' : 'bg-blue-500 hover:bg-blue-600'
              } text-white px-4 py-2 rounded`}
            >
              {isEditMode ? 'Update' : 'Save'}
            </button>
          </div>
        </form>
      )}
    </div>
  );
};

export default CampaignFormPage;
