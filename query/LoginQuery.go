package query

var LoginAdminSQL = `
SELECT
  u."refUserId",
  u."refUserCustId",
  u."refRTId",
  u."refUserFirstName",
  u."refUserLastName",
  rt."refRTName",
  rad."refADHashPass",
  rcd."refCODOPhoneNo1",
  rcd."refCODOEmail"
FROM
  public."Users" u
  JOIN userdomain."refAuthDomain" rad ON rad."refUserId" = u."refUserId"
  JOIN userdomain."refCommunicationDomain" rcd ON rcd."refUserId" = u."refUserId"
  JOIN public."RoleType" rt ON rt."refRTId" = u."refRTId"
WHERE
  rcd."refCODOPhoneNo1" = $1
  AND u."refUserStatus" = true;
`
