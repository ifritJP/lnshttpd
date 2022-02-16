--./test.lns
local _moduleObj = {}
local __mod__ = '@test'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
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
local lnsservlet = _lune.loadModule( 'go/github:com.ifritJP.lnshttpd.src.lns.httpd.lnsservlet' )

local function handle( req )
   local __func__ = '@test.handle'

   print( __func__, req.Url )
   
   local resp = lnsservlet.ResponseInfo.new()
   resp.Header["HogeHeader"] = {"HogeVal"}
   
   resp.StatusCode = 200
   
   resp.Writer = function ( stream )
   
      
      stream:write( "resp1" )
      stream:write( "resp2" )
   end
   
   return resp
end
_moduleObj.handle = handle

local function start( port )

   local servlet = require( "go/github:com.ifritJP.lnshttpd.src.lns.httpd.servlet" )
   
   servlet.Start( port, {lnsservlet.HandlerInfo.new("/code", handle)}, {lnsservlet.HostingInfo.new("./indexableContents", "/contents")} )
end
_moduleObj.start = start

local function __main( argList )

   start( 8080 )
   return 0
end
_moduleObj.__main = __main

do
   local loaded, mess = _lune.loadstring52( [=[
if _lune and _lune._shebang then
  return nil
else
  return arg
end
]=] )
   if loaded ~= nil then
      local args = loaded(  )
      do
         local obj = (args )
         if obj ~= nil then
            local work = obj
            local argList = {""}
            do
               local _exp = work[0]
               if _exp ~= nil then
                  argList[1] = _exp
               end
            end
            for key, val in pairs( work ) do
               if key > 0 then
                  table.insert( argList, val )
               end
            end
            __main( argList )
         else
            -- print( "via lnsc" )
         end
      end
   else
      error( mess )
   end
end

return _moduleObj
