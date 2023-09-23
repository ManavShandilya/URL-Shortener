const BASE_URL = "https://goly-v2xn.onrender.com/";


// {
//     "redirect": "https://www.youtube.com/nerdcademydev",
//     "goly": "short",
//     "random": false
// }

// {
//     "redirect": "https://www.youtube.com/nerdcademydev",
//     "goly": null,
//     "random": true
// }

import axios from "axios";

const API = {
    goly: "/goly",
    golybyid: "/goly/:",// append id while calling
    redirect: "r/" // ridirect
};
const getAPIUrl = (url) => {
    if (Array.isArray(url)) {
        let workUrl = url[0];
        url.shift();
        url = API[workUrl] + url.join("/") + "/";
    } else url = API[url];
    return url;
};
const axioInstance = axios.create({
    baseURL: BASE_URL,
});
axioInstance.interceptors.request.use(async (request) => {
    request.url = getAPIUrl(request.url);
    return request;
});
const axiosFireApi = async (
    url,
    method = "get",
    data = {},
    handleError = false
) => {
    method = method.toLowerCase();
    let allData = {
        url,
        method,
    };
    if (method === "get") {
        allData["params"] = data;
    } else if (method === "post") {
        allData["data"] = data;
    } else if (method === "patch" || method === "put" || method === "delete") {
        let updateId = data.id;
        if (method !== "delete") delete data.id;
        url = [url, updateId];
        allData["data"] = data;
    }
    try {
        const result = await axioInstance(allData);
        return result.data;
    } catch (error) {
        const status = error.response.status;
        if (status === 400 && handleError) handleError(error.response.data);
        console.error(error);
        return null;
    }
};

export { API, getAPIUrl, axioInstance, axiosFireApi, BASE_URL }
