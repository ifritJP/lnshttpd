--lnsservlet.lns
local _moduleObj = {}
local __mod__ = '@lnsservlet'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

if not _lune3 then
   _lune3 = _lune
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


return _moduleObj
