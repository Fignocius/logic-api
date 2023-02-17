DELETE FROM public.logics;

INSERT INTO public.logics (id,"expression",expression_code,created_at,updated_at) VALUES
	 ('f4c01f76-b202-4d13-826b-46cd6fd97249','x OR y','? OR ?','2023-02-16 17:15:00.178','2023-02-16 17:15:00.178'),
	 ('8f1196dc-a2f1-4667-81f2-99023cf7c5ea','x AND y','? AND ?','2023-02-16 17:15:05.170','2023-02-16 17:15:05.170'),
	 ('4998e9c1-8319-4c7f-812f-7fe135fb3eb3','(x AND y) OR z','(? AND ?) OR ?','2023-02-16 17:15:21.382','2023-02-16 17:15:21.382');