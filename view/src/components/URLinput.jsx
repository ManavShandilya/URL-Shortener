/* eslint-disable react/prop-types */
import { useState } from 'react';
import { axiosFireApi } from '../api/config';


function URLinput({ setresponse }) {
    const [url, setUrl] = useState('');
    const [useRandom, setUseRandom] = useState(false);
    const [customShortUrl, setCustomShortUrl] = useState('');

    const handleUrlChange = (event) => {
        setUrl(event.target.value);
    };

    const handleRandomCheckboxChange = () => {
        if (useRandom) {
            setUseRandom(false);
            setCustomShortUrl(''); // Clear customShortUrl when useRandom is true
        } else {
            setUseRandom(true);
        }
    };


    const handleCustomShortUrlChange = (event) => {
        setCustomShortUrl(event.target.value);

    };

    const handleShortenClick = async () => {
        const data = {
            "redirect": url,
            "goly": useRandom ? null : customShortUrl,
            "random": useRandom
        }
        const response = await axiosFireApi("goly", "POST", data);
        setresponse(response);
    };

    return (
        <div className='flex-col mt-8 shadow-xl p-8 container rounded-lg   mx-auto max-w-xl'>
            <div className="flex">
                <input
                    type="text"
                    id="url"
                    name="url"
                    className="w-full p-4 border border-[#284243] rounded-l-lg flex items-center"
                    placeholder="Enter URL to be shortened"
                    value={url}
                    onChange={handleUrlChange}
                />
                <button className='bg-[#284243] text-white p-4 font-semibold rounded-r-lg flex justify-center items-center' onClick={handleShortenClick}>
                    <p className='w-[100px]'>Shorten URL</p>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M8.25 4.5l7.5 7.5-7.5 7.5" />
                    </svg>
                </button>
            </div>
            <div className="flex items-center mt-4">
                <input
                    type="checkbox"
                    id="randomCheckbox"
                    name="randomCheckbox"
                    className="mr-2"
                    checked={useRandom}
                    onChange={handleRandomCheckboxChange}
                />
                <label htmlFor="randomCheckbox">Use Random Short URL</label>
            </div>
            {!useRandom && (
                <input
                    type="text"
                    id="customShortUrl"
                    name="customShortUrl"
                    className="w-full p-4 border border-[#284243] rounded-lg flex items-center mt-4"
                    placeholder="Custom Short URL"
                    value={customShortUrl}
                    onChange={handleCustomShortUrlChange}
                />
            )}
        </div>
    );
}

export default URLinput;
