export const apiUploadFile = (size, uuid, filename) => `/file/upload?size=${size}&uuid=${uuid}&filename=${filename}`; // GET
export const apiDownloadFile = (uuid, filename) => `/file/download?uuid=${uuid}&filename=${filename}`; // GET
export const apiDeleteFile = (uuid, filename) => `/file?uuid=${uuid}&filename=${filename}`; // DELETE
export const apiFileSize = (uuid, filename) => `/file/size?uuid=${uuid}&filename=${filename}`; // GET
