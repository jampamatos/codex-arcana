// File: frontend/master/src/pages/SessionListPage.tsx

import React, { useEffect, useState } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";

// Interface that reflects backend Session model
interface Session {
    id: number;
    title: string;
    date: string;
    location: string;
    notes: string;
    created_at: string;
    updated_at: string;
}

// Initializing states for sessions, loading and error
const SessionListPage: React.FC = () => {
    const { campaignID } = useParams<{ campaignID: string }>();
    const [sessions, setSessions] = useState<Session[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    // Function to fetch all sessions for a specific campaign
    useEffect(() => {
        async function loadSessions() {
            try {
                setLoading(true);
                setError(null);

                const res = await fetch(`http://localhost:3000/api/campaigns/${campaignID}/sessions`);
                if (!res.ok) {
                    throw new Error(`Error fetching sessions: ${res.statusText}`);
                }
                const data: Session[] = await res.json();
                setSessions(data);
            } catch (err: unknown) {
                console.error(err);
                setError(err instanceof Error ? err.message : "An unknown error occurred");
            } finally {
                setLoading(false);
            }
        }
        loadSessions();
    }, [campaignID]);

    if (loading) {
        return (
            <div className="p-4 text-center">
                Loading sessions...
            </div>
        );
    }

    if (error) {
        return (
            <div className="p-4 text-center text-red-500">
                {error}
            </div>
        );
    }

    return (
        <div className="p-4">
            <div className="flex justify-between items-center mb-4">
                <h1 className="text-2xl font-bold">Sessions of {campaignID} Campaign</h1>
                <button
                  onClick={() => navigate(`/campaigns/${campaignID}/sessions/new`)}
                  className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded">
                    Add New Session
                  </button>
            </div>
            { sessions.length === 0 ? (
                <p>No session was found.</p>
            ) : (
                <table className="min-w-full divide-y divide-gray-200 border">
                    <thead className="bg-gray-100">
                        <tr>
                            <th className="px-4 py-2 text-left">Title</th>
                            <th className="px-4 py-2 text-left">Date</th>
                            <th className="px-4 py-2 text-left">Location</th>
                            <th className="px-4 py-2 text-left">Actions</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-200">
                        {sessions.map(s => (
                            <tr key={s.id}>
                                <td className="px-4 py-2">
                                    <Link
                                    to={`/campaigns/${campaignID}/sessions/${s.id}/edit`}
                                    className="text-blue-600 hover:underline">
                                        {s.title}
                                    </Link>
                                </td>
                                <td className="px-4 py-2">
                                    {new Date(s.date).toLocaleDateString()}
                                </td>
                                <td className="px-4 py-2">
                                    {s.location}
                                </td>
                                <td className="px-4 py-2 space-x-2">
                                    <button
                                      onClick={()=> navigate(`/campaigns/${campaignID}/sessions/${s.id}/edit`)}
                                      className="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded">
                                        Edit
                                    </button>
                                    <button
                                      onClick={async () => {
                                        if (!window.confirm('Delete this session?')) return;
                                        const resp = await fetch(
                                            `http://localhost:3000/api/campaigns/${campaignID}/sessions/${s.id}`,
                                            { method: 'DELETE' }
                                        );
                                        if (resp.status === 204) {
                                            setSessions(prev => prev.filter(x => x.id !== s.id));
                                        } else {
                                            alert('Error deleting session');
                                        }
                                      }}
                                      className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded">
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

export default SessionListPage;