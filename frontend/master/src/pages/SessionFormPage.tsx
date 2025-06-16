// File: frontend/master/src/pages/SessionFormPage.tsx

import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

// Interface that reflects the form data structure
interface SessionFormData {
    title: string;
    date: string;
    location: string;
    notes: string;
}

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


const SessionFormPage: React.FC = () => {
    const { campaignID, id } = useParams<{ campaignID: string; id?: string }>();
    const navigate = useNavigate();

    // Form states
    const nowLocal = new Date().toISOString().slice(0, 16); // Format for datetime-local input
    const [formData, setFormData] = useState<SessionFormData>({
        title: "",
        date: nowLocal, // Default to current date and time
        location: "",
        notes: ""
    }); 
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    // Load data in edit mode (when id is present)
    useEffect(() => {
        async function loadSession() {
            if (!id) return; // No need to load if id is not present

            try {
                setLoading(true);
                setError(null);

                const res = await fetch(`http://localhost:3000/api/campaigns/${campaignID}/sessions`);
                if (!res.ok) {
                    throw new Error(`Error fetching session: ${res.statusText}`);
                }
                const data: Session[] = await res.json();
                const sess = data.find(s => s.id === Number(id));
                if (!sess) {
                    throw new Error("Session not found");
                }
                setFormData({
                    title: sess.title,
                    date: sess.date.slice(0, 16), // Format to match datetime-local input
                    location: sess.location,
                    notes: sess.notes
                });
            } catch (err: unknown) {
                console.error(err);
                setError(err instanceof Error ? err.message : "An unknown error occurred");
            } finally {
                setLoading(false);
            }
        }
        loadSession();
    }, [campaignID, id]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try{
            setLoading(true);
            setError(null);

            const url = id
                ? `http://localhost:3000/api/campaigns/${campaignID}/sessions/${id}`
                : `http://localhost:3000/api/campaigns/${campaignID}/sessions`;
            const method = id ? "PUT" : "POST";

            // Prepare payload with RFC3339 date format
            const payload = {
                title: formData.title,
                date: new Date(formData.date).toISOString(),
                location: formData.location,
                notes: formData.notes
            };
            const res = await fetch(url, {
                method,
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload)
            });
            if (!res.ok) {
                throw new Error(`Error saving session: ${res.statusText}`);
            }

            // Redirect to session list page after successful save
            navigate(`/campaigns/${campaignID}/sessions`);
        } catch (err: unknown) {
            console.error(err);
            setError(err instanceof Error ? err.message : "An unknown error occurred");
        } finally {
            setLoading(false);
        }
    }; 

    return (
        <div className="p-4 max-w-md mx-auto">
            <h1 className="text-2xl font-bold mb-4">
                {id ? "Edit Session" : "Create Session"}
            </h1>
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <label className="block mb-1">Title</label>
                    <input 
                      type="text"
                      name="title"
                      value={formData.title}
                      onChange={handleChange}
                      className="w-full border px-2 py-1 rounded" />
                </div>
                <div>
                    <label className="block mb-1">Date</label>
                    <input
                      type="datetime-local"
                      name="date"
                      value={formData.date}
                      onChange={handleChange}
                      className="w-full border px-2 py-1 rounded" />
                </div>
                <div>
                    <label className="block mb-1">Location</label>
                    <input
                      type="text"
                      name="location"
                      value={formData.location}
                      onChange={handleChange}
                      className="w-full border px-2 py-1 rounded" />
                </div>
                <div>
                    <label className="block mb-1">Notes</label>
                    <textarea
                      name="notes"
                      value={formData.notes}
                      onChange={handleChange}
                      className="w-full border px-2 py-1 rounded"
                      rows={4} />
                </div>
                <div className="flex justify-end space-x-2">
                    <button
                      type="button"
                      onClick={() => navigate(-1)}
                      className="px-4 py-2 border rounded"
                    >
                        Cancel
                    </button>
                    <button
                      type="submit"
                      disabled={loading}
                      className="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded"
                    >
                        {loading ? "Saving..." : "Save"}
                    </button>
                </div>
                {error && <p className="text-red-500 mt-2">{error}</p>}
            </form>
        </div>
    );
};

export default SessionFormPage;