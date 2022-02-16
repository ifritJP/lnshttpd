--lns/httpd/bind.lns
local _moduleObj = {}
local __mod__ = '@lns.@httpd.@bind'
local _lune = {}
if _lune6 then
   _lune = _lune6
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
local lnsservlet = _lune.loadModule( 'lns.httpd.lnsservlet' )

local goInStream = {}
setmetatable( goInStream, { ifList = {lnsservlet.inStream,} } )
function goInStream:read( size )

   return nil, ""
end
function goInStream:readAll(  )

   return nil, ""
end
function goInStream.setmeta( obj )
  setmetatable( obj, { __index = goInStream  } )
end
function goInStream.new( reader )
   local obj = {}
   goInStream.setmeta( obj )
   if obj.__init then
      obj:__init( reader )
   end
   return obj
end
function goInStream:__init( reader )

   self.reader = reader
end


local goOutStream = {}
setmetatable( goOutStream, { ifList = {lnsservlet.outStream,} } )
function goOutStream:write( bin )

   return ""
end
function goOutStream.setmeta( obj )
  setmetatable( obj, { __index = goOutStream  } )
end
function goOutStream.new( writer )
   local obj = {}
   goOutStream.setmeta( obj )
   if obj.__init then
      obj:__init( writer )
   end
   return obj
end
function goOutStream:__init( writer )

   self.writer = writer
end


do
   local _
   local _1 = goInStream.new(0)
   local _2 = goOutStream.new(0)
end


return _moduleObj
