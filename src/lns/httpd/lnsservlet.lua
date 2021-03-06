--lns/httpd/lnsservlet.lns
local _moduleObj = {}
local __mod__ = '@lns.@httpd.@lnsservlet'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune6 then
   _lune6 = _lune
end
local inStream = {}
_moduleObj.inStream = inStream
function inStream.setmeta( obj )
  setmetatable( obj, { __index = inStream  } )
end
function inStream.new(  )
   local obj = {}
   inStream.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function inStream:__init(  )

end


local outStream = {}
_moduleObj.outStream = outStream
function outStream.setmeta( obj )
  setmetatable( obj, { __index = outStream  } )
end
function outStream.new(  )
   local obj = {}
   outStream.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function outStream:__init(  )

end


local luaInStream = {}
setmetatable( luaInStream, { ifList = {inStream,} } )
_moduleObj.luaInStream = luaInStream
function luaInStream:readStream( mode )

   do
      local bin = self.stream:read( mode )
      if bin ~= nil then
         return bin, ""
      end
   end
   
   return nil, "err"
end
function luaInStream:read( size )

   return self:readStream( size )
end
function luaInStream:readAll(  )

   return self:readStream( "*a" )
end
function luaInStream.setmeta( obj )
  setmetatable( obj, { __index = luaInStream  } )
end
function luaInStream.new( stream )
   local obj = {}
   luaInStream.setmeta( obj )
   if obj.__init then
      obj:__init( stream )
   end
   return obj
end
function luaInStream:__init( stream )

   self.stream = stream
end


local luaOutStream = {}
setmetatable( luaOutStream, { ifList = {outStream,} } )
_moduleObj.luaOutStream = luaOutStream
function luaOutStream:write( bin )

   local _
   local _1, err = self.stream:write( bin )
   if err ~= nil then
      return err
   end
   
   return ""
end
function luaOutStream.setmeta( obj )
  setmetatable( obj, { __index = luaOutStream  } )
end
function luaOutStream.new( stream )
   local obj = {}
   luaOutStream.setmeta( obj )
   if obj.__init then
      obj:__init( stream )
   end
   return obj
end
function luaOutStream:__init( stream )

   self.stream = stream
end




local ResponseInfo = {}
_moduleObj.ResponseInfo = ResponseInfo
function ResponseInfo.new(  )
   local obj = {}
   ResponseInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function ResponseInfo:__init() 
   self.StatusCode = 200
   self.Header = {}
   self.Writer = nil
   self.Txt = ""
end
function ResponseInfo.setmeta( obj )
  setmetatable( obj, { __index = ResponseInfo  } )
end


local RequestInfo = {}
_moduleObj.RequestInfo = RequestInfo
function RequestInfo.setmeta( obj )
  setmetatable( obj, { __index = RequestInfo  } )
end
function RequestInfo.new( Method, Url, Header, Reader )
   local obj = {}
   RequestInfo.setmeta( obj )
   if obj.__init then
      obj:__init( Method, Url, Header, Reader )
   end
   return obj
end
function RequestInfo:__init( Method, Url, Header, Reader )

   self.Method = Method
   self.Url = Url
   self.Header = Header
   self.Reader = Reader
end



local HandlerInfo = {}
_moduleObj.HandlerInfo = HandlerInfo
function HandlerInfo.setmeta( obj )
  setmetatable( obj, { __index = HandlerInfo  } )
end
function HandlerInfo.new( Path, Handler )
   local obj = {}
   HandlerInfo.setmeta( obj )
   if obj.__init then
      obj:__init( Path, Handler )
   end
   return obj
end
function HandlerInfo:__init( Path, Handler )

   self.Path = Path
   self.Handler = Handler
end


local HostingInfo = {}
_moduleObj.HostingInfo = HostingInfo
function HostingInfo.setmeta( obj )
  setmetatable( obj, { __index = HostingInfo  } )
end
function HostingInfo.new( LocalPath, UrlPath )
   local obj = {}
   HostingInfo.setmeta( obj )
   if obj.__init then
      obj:__init( LocalPath, UrlPath )
   end
   return obj
end
function HostingInfo:__init( LocalPath, UrlPath )

   self.LocalPath = LocalPath
   self.UrlPath = UrlPath
end


return _moduleObj
