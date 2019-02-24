grammar Path;

query
  : AND? andExpr AND?
  | AND? orExpr AND?
  | AND? tag AND?
  | AND? anyExpr AND?
  |
  ;

anyExpr
  : ANY
  | ANY (AND ANY)+
  ;

andExpr
  : orExpr (AND orExpr)+
  | orExpr (AND tag)+
  | tag (AND orExpr)+
  | tag (AND tag)+
  ;

orExpr
  : tag (OR tag)+
  ;

tag
  : classifierTagExpr
  | nameTagExpr
  ;

classifierTagExpr
  : classifier=STRING op=EQ name=ANY
  ;

nameTagExpr
  : name=STRING
  | classifier=STRING op=EQ name=STRING
  ;

AND : '/' ;
ANY : ('*' | '**' ) ;
EQ : '=' ;
OR : ',' ;
STRING : ~('/' | ',' | '=')+ ;
