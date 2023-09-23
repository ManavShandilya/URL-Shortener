/* eslint-disable react/prop-types */
import { useState } from "react";
import { API, BASE_URL } from "../api/config";

function ShortenResponse({ response }) {
    const [copied, setCopied] = useState(false);

    const handleCopyClick = () => {
        const copyText = BASE_URL + API.redirect + response?.data?.goly;
        navigator.clipboard.writeText(copyText)
            .then(() => {
                setCopied(true);
                setTimeout(() => setCopied(false), 1500); // Reset "copied" state after a short delay
            })
            .catch((error) => {
                console.error("Copy failed:", error);
            });
    };

    return (
        response ?
            <div className="container mx-auto bg-[#d1ecf1] p-6 flex justify-between items-center mt-8 max-w-xl shadow-xl rounded-lg  flex-col gap-6">
                <div className="flex justify-center items-center gap-8">
                    <div className=" text-[#284243] font-bold text-lg md:text-xl">
                        <a href={BASE_URL + API.redirect + response?.data?.goly}>{BASE_URL + response?.data?.goly}</a>
                    </div>
                    <div className='cursor-pointer' onClick={handleCopyClick}>
                        {copied ? (
                            <span className="text-[#3fed8a]">Copied!</span>
                        ) : (
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6 text-[#284243]">
                                <path d="M7.5 3.375c0-1.036.84-1.875 1.875-1.875h.375a3.75 3.75 0 013.75 3.75v1.875C13.5 8.161 14.34 9 15.375 9h1.875A3.75 3.75 0 0121 12.75v3.375C21 17.16 20.16 18 19.125 18h-9.75A1.875 1.875 0 017.5 16.125V3.375z" />
                                <path d="M15 5.25a5.23 5.23 0 00-1.279-3.434 9.768 9.768 0 016.963 6.963A5.23 5.23 0 0017.25 7.5h-1.875A.375.375 0 0115 7.125V5.25zM4.875 6H6v10.125A3.375 3.375 0 009.375 19.5H16.5v1.125c0 1.035-.84 1.875-1.875 1.875h-9.75A1.875 1.875 0 013 20.625V7.875C3 6.839 3.84 6 4.875 6z" />
                            </svg>
                        )}
                    </div>
                </div>
                {/* <div className="text-[#284243] font-semibold mr-8">
                    <div className="flex space-x-2">
                        <button className="px-4 py-2 bg-white text-[#17a2b8] rounded-md hover:bg-[#3fed8a] hover:text-white">
                            Update
                        </button>
                        <button className="px-4 py-2 bg-white text-[#e74c3c] rounded-md hover:bg-[#e74c3c] hover:text-white">
                            Delete
                        </button>
                    </div>
                </div> */}
            </div>
            : null
    );
}

export default ShortenResponse;
