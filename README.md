# localcache
Local cache for storing any data e.g. configuration

Very often, a service receives a configuration from an external server using, for example, the WEB API.
If you are creating a fault tolerant application e.g. SCADA, your application shall work regardless of the current state of the SCADA WEB API configuration server. The main idea of this library is to save the configuration to a local cache file and restore the configuration from the file in case the WEB API configuration server fails.

## How to use

```Go
cache := localcache.New(cachePath)

if loadFromCacheOnly {
	return cache.Load()
}
  
 profileData, err = api.GetProfileFromWebApi()
 if err != nil {
 	return cache.Load()
}
   
cache.Save(profileData)
return profileData
```
Before saving, this library checks the checksum of the data and the saved file, and saves data only if the checksums differ.
