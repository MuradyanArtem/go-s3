export const apiUploadFile = (size, uuid, filename) => `/api/file/upload?size=${size}&uuid=${uuid}&filename=${filename}`; // GET
export const apiDownloadFile = (uuid, filename) => `/api/file/download?uuid=${uuid}&filename=${filename}`; // GET
export const apiDeleteFile = (uuid, filename) => `/api/file?uuid=${uuid}&filename=${filename}`; // DELETE
export const apiFileSize = (uuid, filename) => `/api/file/size?uuid=${uuid}&filename=${filename}`; // GET
