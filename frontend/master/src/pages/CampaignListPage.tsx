// File: frontend/master/src/pages/CampaignListPage.tsx

import React, { useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";

// Interface that reflects backend Campaign model
interface Campaign {
    id: number;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
}

const CampaignListPage: React.FC = () => {
    const [campaigns, setCampaigns] = useState<Campaign[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const navigate = useNavigate();

    // Function to fetch all backend campaigns
    const fetchCampaigns = async () => {
        try {
            setLoading(true);
            setError(null);

            const response = await fetch('http://localhost:3000/api/campaigns');
            if (!response.ok) {
                throw new Error(`Error fetching campaigns: ${response.statusText}`);
            }

            const data: Campaign[] = await response.json();
            setCampaigns(data);
        } catch (err: unknown) {
            console.error(err);
            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError("An unknown error occurred");
            }
        } finally {
            setLoading(false);
        }
    };

    // Function to delete a campaign
    const deleteCampaign = async (id: number) => {
        if (!window.confirm("Are you sure you want to delete this campaign?")) {
            return;
        }

        try {
            const response = await fetch(`http://localhost:3000/api/campaigns/${id}`, { method: 'DELETE'});
            if (response.status === 204) {
                // Locally remove the campaign from the list
                setCampaigns(prev => prev.filter(c => c.id !== id));
            } else if (response.status === 404) {
                alert("Campaign not found.");
                // Maybe refresh the list
                fetchCampaigns();
            } else {
                throw new Error(`Error deleting campaign: ${response.statusText}`);
            }
        } catch (err: unknown) {
            console.error(err);
            if (err instanceof Error) {
                alert(err.message || "An error occurred while deleting the campaign.");
            } else {
                alert("An error occurred while deleting the campaign.");
            }
        }
    };

    useEffect(() => {
        fetchCampaigns();
    }, [])

    if (loading) {
        return (
            <div className="p-4 text-center">
                <p>Loading campaigns...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="p-4 text-center text-red-500">
                <p>{error}</p>
            </div>
        );
    }

    return (
        <div className="p-4">
            <div className="flex justify-between items-center mb-4">
                <h1 className="text-2xl font-bold">Campaigns</h1>
                <button
                    onClick={() => navigate('/campaigns/new')}
                    className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
                    >
                        New Campaign
                    </button>
            </div>

            {campaigns.length === 0 ? (
                <p>No campaign found.</p>
            ) : (
                <table className="min-w-full divide-y divide-gray-200 border">
                    <thead className="bg-gray-100">
                        <tr>
                            <th className="px-4 py-2 text-left">Name</th>
                            <th className="px-4 py-2 text-left">Description</th>
                            <th className="px-4 py-2 text-left">Actions</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-200">
                        {campaigns.map(c => (
                            <tr key={c.id}>
                                <td className="px-4 py-2">
                                    <Link to={`/campaigns/${c.id}/sessions`} className="text-blue-600 hover:underline">
                                        {c.name}
                                    </Link>
                                </td>
                                <td className="px-4 py-2">{c.description}</td>
                                <td className="px-4 py-2 text-center space-x-2">
                                    <button
                                        onClick={() => navigate(`/campaigns/${c.id}/edit`)}
                                        className="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded"
                                    >
                                        Edit
                                    </button>
                                    <button
                                        onClick={() => deleteCampaign(c.id)}
                                        className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
                                    >
                                        Delete
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    );

};

export default CampaignListPage;